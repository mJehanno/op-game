
export interface DifficultyLevelData {
    order: number;
    value: string;
}

export enum DifficultyLevel {
    Easy = "easy",
    Medium = "medium",
    Hard = "hard"
}

export function cmpDifficulty(a: DifficultyLevel, b: DifficultyLevel): number {
    if (a == b) {
        return 0
    }
    switch (a) {
        case DifficultyLevel.Easy:
            return -1;   
        case DifficultyLevel.Hard:
            return 1;
        case DifficultyLevel.Medium:
            return b == DifficultyLevel.Easy ? 1 : -1
    }
}

export enum GameMode {
    Mult= "mult",
    Add = "add",
    Min = "min",
    Divid = "divid"
}

export interface GameState {
    endingDialogVisible : boolean;
    maxSec: number;
    maxLife: number;
    currentLife: number;
}