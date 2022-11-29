import { Component, OnInit } from '@angular/core';
import {ProjectServices} from "../services/project.services";
import {IProj} from "../models/proj.model";

@Component({
  selector: 'app-project-page',
  templateUrl: './project-page.component.html',
  styleUrls: ['./project-page.component.css']
})
export class ProjectPageComponent implements OnInit {
  projects: IProj[] = []
  loading = false
  searchNameStr = ''

  constructor(private projectService: ProjectServices) {
  }

  ngOnInit(): void {
    this.loading = true
    this.projectService.getAll().subscribe(projects => {
      console.log(projects)
      this.projects = projects.data
      this.loading = false
    })
  }

  addMyProjectToDB(name: String) {
    //to-do
  }
}
