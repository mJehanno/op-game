package update

import (
	"context"
	"fmt"
	"io"
	"mult-game/backend/logger"
	"mult-game/backend/version"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/blang/semver"
	"github.com/google/go-github/v59/github"
	"go.uber.org/fx"
	"golang.org/x/exp/slices"
)

const (
	owner = "mJehanno"
	repo  = "op-game"
)

type ErrorHandler struct {
	log *logger.Logger
}

func NewErrorHandler(log *logger.Logger) *ErrorHandler {
	return &ErrorHandler{
		log: log,
	}
}

func (eh *ErrorHandler) handleErr(err error, message string) {
	if err != nil {
		eh.log.ErrLogger.WithError(err).Error(message)
	}
}

type Updater struct {
	log      *logger.Logger
	eh       *ErrorHandler
	vm       *version.VersionManager
	client   *github.Client
	latest   semver.Version
	rel      *github.RepositoryRelease
	platform string
	dlPath   string
}

type UpdaterParams struct {
	fx.In

	Log    *logger.Logger
	Eh     *ErrorHandler
	Vm     *version.VersionManager
	Client *github.Client
}

func NewUpdater(p UpdaterParams) *Updater {
	platf := fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
	return &Updater{
		log:      p.Log,
		eh:       p.Eh,
		vm:       p.Vm,
		client:   p.Client,
		platform: platf,
	}
}

func (u *Updater) isCurrentLatest(ctx context.Context) bool {
	rel, _, err := u.client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		u.eh.handleErr(err, "failed to retrieve latest release")
		return true
	}
	u.rel = rel

	latest := semver.MustParse(strings.ReplaceAll(rel.GetTagName(), "v", ""))
	u.latest = latest
	return latest.EQ(semver.MustParse(u.vm.GetCurrentVersion()))
}

func (u *Updater) getAsset() (*github.ReleaseAsset, error) {

	index := slices.IndexFunc(u.rel.Assets, func(ra *github.ReleaseAsset) bool {
		return strings.Contains(*ra.Name, u.platform)
	})

	if index == -1 {
		err := fmt.Errorf("asszt not found")
		u.log.ErrLogger.WithError(err).Error("failed to retrieve corresponding release asset")
		return nil, err
	}

	return u.rel.Assets[index], nil
}

func (u *Updater) downloadAsset(ctx context.Context, asset github.ReleaseAsset) error {
	reader, redirect, err := u.client.Repositories.DownloadReleaseAsset(ctx, owner, repo, asset.GetID(), http.DefaultClient)
	if err != nil {
		err = fmt.Errorf("failed to download release asset -> %w", err)
		return err
	}

	if redirect != "" {
		return fmt.Errorf("need to learn how to use the redirect url")
	}

	u.dlPath = path.Join(os.TempDir(), *asset.Name)

	f, err := os.Create(u.dlPath)
	if err != nil {
		err = fmt.Errorf("failed to create temp downloaded release asset -> %w", err)
		return err
	}

	defer func() {
		f.Close()
		reader.Close()
	}()

	_, err = io.Copy(f, reader)
	if err != nil {
		err = fmt.Errorf("failed to write downloaded release asset -> %w", err)
		return err
	}
	return nil
}

func (u *Updater) installNewRelease() error {
	exePath, err := os.Executable()
	if err != nil {
		u.log.ErrLogger.WithError(err).Error("failed to retrieve current process executable")
		return err
	}

	err = os.Rename(exePath, fmt.Sprintf("%s-old", exePath))
	if err != nil {
		u.log.ErrLogger.WithError(err).Error("failed to rename old binary")
		return err
	}

	if strings.Contains(u.platform, "linux") {
		os.Rename(u.dlPath, exePath)
		os.Chmod(exePath, 0775)
		err := exec.Command(exePath).Run()
		if err != nil {
			u.log.ErrLogger.WithError(err).Error("failed to launch new binary")
			// rollingback
			err = os.Remove(exePath)
			if err != nil {
				u.log.ErrLogger.WithError(err).Error("failed to remove new binary")
			}
			err = os.Rename(fmt.Sprintf("%s-old", exePath), exePath)
			if err != nil {
				u.log.ErrLogger.WithError(err).Error("failed to rename old binary")
			}
			return err
		}
	}

	return nil
}

func (u *Updater) DoSelfUpdate(ctx context.Context) {
	if u.isCurrentLatest(ctx) {
		u.log.DebugLogger.Debug("current version is the latest")
		return
	}
	u.log.DebugLogger.Debugf("latest version is %s, current is %s", u.latest, u.vm.GetCurrentVersion())

	asset, err := u.getAsset()
	u.eh.handleErr(err, "failed to find asset")
	if err != nil {
		return
	}

	err = u.downloadAsset(ctx, *asset)
	u.eh.handleErr(err, "failed to download asset")
	if err != nil {
		return
	}

	err = u.installNewRelease()
	u.eh.handleErr(err, "failed to install new release")
	if err == nil {
		os.Exit(0)
	}

}
