export interface Game {
    state: State;
}

export enum State {
    Ready,
    Started,
    Over
}
