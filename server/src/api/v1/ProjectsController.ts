import {Get, Post, JsonController, Param, Body, Req, UseBefore} from 'routing-controllers';
import {JSONWebToken} from '../../utils/JSONWebToken';
import {Project} from '../../entities/project.model';
import {UserAuthenticatorMiddleware} from '../../middleware/UserAuthenticatorMiddleware';
import {CORSMiddleware} from '../../middleware/CORSMiddleware';
import {LoggerFactory} from '../../utils/LoggerFactory';
import {Service} from 'typedi';

@JsonController('/projects')
@UseBefore(UserAuthenticatorMiddleware, CORSMiddleware)
@Service()
export class ProjectsController {
    public constructor(private loggerFactory: LoggerFactory) { }

    @Get('/voter/:id')
    public getProjectsByVoterID(@Param('id') voterId: string, @Req() request: any): any {
        let enrollmentID = new JSONWebToken(request).getUserID();

        return request.blockchain.query('getProjectsForVoter', [voterId], enrollmentID);
    }

    @Post('/')
    public post(@Body() project: Project, @Req() request: any): any {
        let enrollmentID = new JSONWebToken(request).getUserID();
        let projectString = [JSON.stringify(project)];
        console.log('projectString = ' + projectString);
        console.log('project = ' + project);
        console.log('toJSON = ' + project.toJSON());
        for (let i in project) {
            if (project.hasOwnProperty(i)) {
                console.log('pr...i = ' + i + ' => ' + project[i]);
            }
        }
        for (let i in project.toJSON()) {
            if (project.toJSON().hasOwnProperty(i)) {
                console.log('pr.toJSON..i = ' + i + ' => ' + project.toJSON()[i]);
            }
        }
        return request.blockchain.invoke('addProject', [JSON.stringify(project)], enrollmentID);
    }
}
