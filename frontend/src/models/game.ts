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
    err?: string; 
    streak: number;
    endingDialogVisible : boolean;
    level: DifficultyLevel;
    currentlife: number;
    user? :  string;
    answer?: number;
}