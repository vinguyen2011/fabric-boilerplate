'use strict';

export class Vote {

    public constructor(
                        private _voterId: string,
                        private _projectId: string,
                        private _votePercent: number
                       ) {
    }

    public get voterId(): string {
        return this._voterId;
    }

    public get projectId(): string {
        return this._projectId;
    }

    public get votePercent(): number {
        return this._votePercent;
    }

    public toJSON(): any {
        return {
            'voterId': this.voterId,
            'projectId': this.projectId,
            'votePercent': this.votePercent
        };
    }
}
