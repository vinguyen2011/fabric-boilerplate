'use strict';

import * as shortID from 'shortid';

export class Project {
    private _projectID: string;

    public constructor(private _description: string,
                       private _location: string,
                       private _tags: string[],
                       private _voteRestriction: string,
                       private _expiryDate: number,
                       private _cost: number,
                       private _costCovered: number
                       ) {
        this._projectID = shortID.generate();
    }

    public get projectID(): string {
        return this._projectID;
    }

    public get description(): string {
        return this._description;
    }

    public get location(): string {
        return this._location;
    }

    public get tags(): string[] {
        return this._tags;
    }

    public get voteRestriction(): string {
        return this._voteRestriction;
    }

    public get expiryDate(): number {
        return this._expiryDate;
    }

    public get cost(): number {
        return this._cost;
    }

    public get costCovered(): number {
        return this._costCovered;
    }

    public toJSON(): any {
        return {
            'projectId': this.projectID,
            'description': this.description,
            'location': this.location,
            'tags': this.tags,
            'voteRestriction': this.voteRestriction,
            'expiryDate': this.expiryDate,
            'cost': this.cost,
            'costCovered': this.costCovered
        };
    }
}
