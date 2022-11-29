import {Injectable} from "@angular/core";
import {IProj} from "../models/proj.model";

@Injectable({
  providedIn: 'root'
})
export class MyprojectServices {

  myProjects: IProj[] = []

  constructor() {
  }

  getMyProjects() {
    return this.myProjects;
  }

  addMyProject(project: IProj) {
    this.myProjects.push(project);
  }

  deleteMyProject(project: IProj){
    this.myProjects.pop()
  }


}
