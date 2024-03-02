export enum DifficultyLevel {
    Easy = "easy",
    Medium = "medium",
    Hard = "hard"
}

export interface GamePrompt {
    x: number;
    y: number;
    result: number;
    prompt: string;
}

export interface GameState {
    streak: number;
    endingDialogVisible : boolean;
    maxSec: number;
    maxLife: number;
    currentLife: number;
    user? :  string;
}