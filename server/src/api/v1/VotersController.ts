import {Get, Post, JsonController, Param, Body, Req, UseBefore} from 'routing-controllers';
import {JSONWebToken} from '../../utils/JSONWebToken';
import {Voter} from '../../entities/voter.model';
import {Vote} from '../../entities/vote.model';
import {UserAuthenticatorMiddleware} from '../../middleware/UserAuthenticatorMiddleware';
import {CORSMiddleware} from '../../middleware/CORSMiddleware';
import {LoggerFactory} from '../../utils/LoggerFactory';
import {Service} from 'typedi';

@JsonController('/voter')
@UseBefore(UserAuthenticatorMiddleware, CORSMiddleware)
@Service()
export class VotersController {
    public constructor(private loggerFactory: LoggerFactory) { }

    @Get('/:id')
    public getVotersByVoterID(@Param('id') voterId: string, @Req() request: any): any {
        let enrollmentID = new JSONWebToken(request).getUserID();
        return request.blockchain.query('getVoter', [voterId], enrollmentID);
    }

    @Post('/')
    public postNewVoter(@Body() voter: Voter, @Req() request: any): any {
        let enrollmentID = new JSONWebToken(request).getUserID();
        return request.blockchain.invoke('addVoter', [JSON.stringify(voter)], enrollmentID);
    }

    @Post('/vote')
    public postVote(@Body() vote: Vote, @Req() request: any): any {
        let enrollmentID = new JSONWebToken(request).getUserID();
        return request.blockchain.invoke('vote', [JSON.stringify(vote)], enrollmentID);
    }

    @Get('/voteForProject/:id')
    public getVoteForProjectByVoter(@Param('id') projectId: string, @Req() request: any): any {
        let enrollmentID = new JSONWebToken(request).getUserID();
        //assume user id = voter id
        let voterId = enrollmentID;
        return request.blockchain.query('getVoteForProjectByVoter', [projectId, voterId], enrollmentID);
    }

}
