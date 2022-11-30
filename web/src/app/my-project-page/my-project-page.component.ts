import { Component, OnInit } from '@angular/core';
import {DatabaseProjectServices} from "../services/database-project.services";
import {IProj} from "../models/proj.model";

@Component({
  selector: 'app-my-project-page',
  templateUrl: './my-project-page.component.html',
  styleUrls: ['./my-project-page.component.css']
})
export class MyProjectPageComponent implements OnInit {
  myProjects: IProj[] = []

  loading = false
  constructor(private myProjectService: DatabaseProjectServices) { }

  ngOnInit(): void {
    this.loading = true
    this.myProjectService.getAll().subscribe(projects => {
      console.log(projects)
      this.myProjects = projects.data
      this.loading = false
    })
  }

}
