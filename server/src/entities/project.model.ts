'use strict';

import * as shortID from 'shortid';

export class Project {
    private _projectID: string;

    public constructor(private _description: string,
                       private _name: string,
                       private _location: string,
                       private _tags: string[],
                       private _voteRestrictionField: string,
                       private _voteRestrictionValues: string[],
                       private _expiryDate: number,
                       private _cost: number,
                       private _costCovered: number,
                       private _pictureId: number
                       ) {
        this._projectID = shortID.generate();
    }

    public get projectID(): string {
        return this._projectID;
    }

    public get name(): string {
        return this._name;
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

    public get voteRestrictionField(): string {
        return this._voteRestrictionField;
    }

    public get voteRestrictionValues(): string[] {
        return this._voteRestrictionValues;
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

    public get pictureId(): number {
        return this._pictureId;
    }

    public toJSON(): any {
        return {
            'projectId': this.projectID,
            'name': this.name,
            'description': this.description,
            'location': this.location,
            'tags': this.tags,
            'voteRestrictionField': this.voteRestrictionField,
            'voteRestrictionValues': this.voteRestrictionValues,
            'expiryDate': this.expiryDate,
            'cost': this.cost,
            'costCovered': this.costCovered,
            'pictureId': this.pictureId
        };
    }
}
