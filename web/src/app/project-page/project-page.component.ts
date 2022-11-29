import { Component, OnInit } from '@angular/core';
import {IRequest} from "../models/request.model";
import {ProjectServices} from "../services/project.services";
import {MyprojectServices} from "../services/myproject.services";
import {IProj} from "../models/proj.model";

@Component({
  selector: 'app-project-page',
  templateUrl: './project-page.component.html',
  styleUrls: ['./project-page.component.css']
})
export class ProjectPageComponent implements OnInit {
  requestData: IRequest
  projects: IProj[] = []
  myProjects: IProj[] = []
  loading = false

  constructor(private projectService: ProjectServices,
              private myprojectService: MyprojectServices) {
  }

  ngOnInit(): void {
    this.loading = true
    this.projectService.getAll().subscribe(projects => {
      console.log(projects)
      this.requestData = projects
      this.projects = this.requestData.data
      /*this.projects = projects.data*/
      this.loading = false
    })
  }

  addMyProject(project: IProj) {
    this.myprojectService.addMyProject(project)
    this.myProjects.push(project)
    console.log(this.myProjects)
  }
}
