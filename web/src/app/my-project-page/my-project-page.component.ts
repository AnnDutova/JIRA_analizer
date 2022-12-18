import { Component, OnInit } from '@angular/core';
import {DatabaseProjectServices} from "../services/database-project.services";
import {IProj} from "../models/proj.model";
import {CheckedProject} from "../models/check-element.model";

@Component({
  selector: 'app-my-project-page',
  templateUrl: './my-project-page.component.html',
  styleUrls: ['./my-project-page.component.css']
})
export class MyProjectPageComponent implements OnInit {
  myProjects: IProj[] = []
  checked: Map<any, any> = new Map();
  loading = false
  noProjects: boolean = false
  inited: boolean = false

  constructor(private myProjectService: DatabaseProjectServices) { }

  ngOnInit(): void {
    this.loading = true
    this.myProjectService.getAll().subscribe(projects => {
      this.noProjects = projects.data.length == 0;
      console.log(projects)
      this.myProjects = projects.data
      this.loading = false
      this.inited = true
    }, error => {
      if (error.status == 0){
        alert("Unable to connect to backend")
      }
      if (error.status == 400){
        alert("Unable to connect to DB")
      }
    })
  }

  childOnChecked(project: CheckedProject){
    if (project.Checked) {
      this.checked.set(project.Name, project.Id)
    }else if (this.checked.has(project.Name)){
      this.checked.delete(project.Name)
    }
    console.log("Parent ", project.Name, project.Id)
  }

}
