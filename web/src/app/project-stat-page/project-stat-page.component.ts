import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {DatabaseProjectServices} from "../services/database-project.services";

@Component({
  selector: 'project-stat-page',
  templateUrl: './project-stat-page.component.html',
  styleUrls: ['./project-stat-page.component.css']
})
export class ProjectStatPageComponent implements OnInit {
  projects: string[] = []
  ids: string[] = []

  constructor(private route: ActivatedRoute, private projectService: DatabaseProjectServices) {
    this.projects = this.route.snapshot.queryParamMap.getAll("keys")
    this.ids = this.route.snapshot.queryParamMap.getAll("value")
  }

  ngOnInit(): void {
    console.log(this.projects, this.ids)
  }

}
