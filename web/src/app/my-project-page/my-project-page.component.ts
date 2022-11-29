import { Component, OnInit } from '@angular/core';
import {MyprojectServices} from "../services/myproject.services";
import {IProj} from "../models/proj.model";

@Component({
  selector: 'app-my-project-page',
  templateUrl: './my-project-page.component.html',
  styleUrls: ['./my-project-page.component.css']
})
export class MyProjectPageComponent implements OnInit {
  myProjects: IProj[] = []

  loading = false
  constructor(private myprojectService: MyprojectServices) { }

  ngOnInit(): void {
    this.loading = true
    this.myProjects = this.myprojectService.getMyProjects();
    this.loading = false
  }

}
