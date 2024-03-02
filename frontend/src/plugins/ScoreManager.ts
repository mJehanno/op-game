import {score} from '../../wailsjs/go/models'
import { AddScore, GetScore } from '../../wailsjs/go/score/ScoreService';
import  { type App, type Plugin } from 'vue';
import type { ScoreManager } from '@/models/ScoreManager';


export const ScoreManagerService: Plugin = {
    install(app: App, options: ScoreManager) {
            app.config.globalProperties.$getScore = (game: string, level: string): Promise<score.Score[]> => {
                if (!options) {
                    return GetScore(game, level);
                }
                return options.getScore(game,level);
            } 
            app.config.globalProperties.$addScore = (sc: score.Score): Promise<void> => {
                if (!options) {
                    return AddScore(sc);
                }
                return options.addScore(sc);
            }

            app.provide<ScoreManager>("score", {addScore: app.config.globalProperties.$addScore, getScore: app.config.globalProperties.$getScore})
    }
}