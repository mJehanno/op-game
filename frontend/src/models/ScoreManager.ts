import type { score } from "../../wailsjs/go/models";

export type GetScoreFunc = (game: string,difficulty: string) => Promise<score.Score[]>
export type AddScoreFunc = (sc: score.Score) => Promise<void>

export interface ScoreManager{
    getScore: GetScoreFunc;
    addScore: AddScoreFunc;
}