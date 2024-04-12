import { defineStore } from "pinia";
import { useGameInfosStore } from "./game-infos";
import {score} from '../../wailsjs/go/models';
import {computed, inject, reactive} from 'vue';
import type { ScoreManager } from "@/models/ScoreManager";
import { GameMode } from "@/models/game";
import { DifficultyLevel, cmpDifficulty } from "@/models/game";
import { forIn } from "lodash";

export const useScoresStore = defineStore("scores", () => {
    const scoreManager = inject("score") as ScoreManager;
    const gameInfos = useGameInfosStore();
    const s = new score.Score()
    const current = reactive({id: -1,username: 'xxx', score: 0, difficulty: gameInfos.difficultyLevel, game: gameInfos.selectedGame, created_at: new Date()} as score.Score);
    // current need to be set to original value at the end of the game
    const scores = reactive(new Map<GameMode, Map<DifficultyLevel,score.Score[]>>());
    const sortedScore = computed(() => sortMap(scores))

    function resetCurrent() {
        current.score = 0;
        current.username= 'xxx';
    }

    function increaseCurrent() {
        current.score ++;
    }

    function setCurrentScoreUsername(name:string) {
        current.username = name;
    }

    async function saveScore() {
        try {
            await scoreManager.addScore(current);
            resetCurrent()
        } catch (error) {
            return error;
        }
    }

    function sortMap(m :Map<GameMode, Map<DifficultyLevel, score.Score[]>>):Map<GameMode, Map<DifficultyLevel, score.Score[]>> {
        const res = new Map<GameMode, Map<DifficultyLevel, score.Score[]>>();
        for(let inner of m.entries()) {
            res.set(inner[0], new Map([...inner[1]].sort((a,b) => {
                return cmpDifficulty(a[0], b[0]);
            })));
        }  

        return res;
    }

    async function getScoreByGameAndLevel(game: GameMode, level: DifficultyLevel){
        try {
            const dbScores = await scoreManager.getScore(game, level);
            var inner :Map<DifficultyLevel, score.Score[]>;
            if (!scores.has(game)) {
                inner = new Map<DifficultyLevel, score.Score[]>();
            }else {
                inner = scores.get(game) as Map<DifficultyLevel, score.Score[]>;
            }
            inner.set(level, dbScores);
            scores.set(game, inner);
        } catch (err) {
            return err;
        }
    }

    async function getAllScoresByGame(game: GameMode) {
        forIn(DifficultyLevel, async (value: string, _: string) => {
            try {
                let scs = await scoreManager.getScore(game,value);
                if (!scs) {
                    scs = [];
                }

                var inner: Map<DifficultyLevel, score.Score[]>;
                if (!scores.has(game)) {
                    inner = new Map<DifficultyLevel, score.Score[]>();
                }else {
                    inner = scores.get(game) as Map<DifficultyLevel, score.Score[]>;
                }
                
                inner.set(value as DifficultyLevel, scs);
                scores.set(game, inner);
            } catch (error) {
                return error
            }
        }) 
    }

    return {current, scores, sortedScore, increaseCurrent, setCurrentScoreUsername, saveScore, getScoreByGameAndLevel, getAllScoresByGame}
})

