'use strict';

export class Voter {

    private _projectIds: string[];

    public constructor(
                        private _voterId: string,
                        private _location: string,
                        private _gender: string,
                        private _dob: number
                       ) {
        this._projectIds = [];
    }

    public get voterId(): string {
        return this._voterId;
    }

    public get location(): string {
        return this._location;
    }

    public get gender(): string {
        return this._gender;
    }

    public get dob(): number {
        return this._dob;
    }

    public get projectIds(): string[] {
        return this._projectIds;
    }

    public toJSON(): any {
        return {
            'voterId': this.voterId,
            'location': this.location,
            'gender': this.gender,
            'dob': this.dob,
            'projectIds': this.projectIds
        };
    }
}
