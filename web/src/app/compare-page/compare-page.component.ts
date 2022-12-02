import { Component, OnInit } from '@angular/core';
import {IProj} from "../models/proj.model";
import {DatabaseProjectServices} from "../services/database-project.services"
import { Router } from '@angular/router';
import {CheckedProject} from "../models/check-element.model";


@Component({
  selector: 'app-compare-page',
  templateUrl: './compare-page.component.html',
  styleUrls: ['./compare-page.component.css']
})
export class ComparePageComponent implements OnInit {
  projects: IProj[] = []
  checked: Map<any, any> = new Map();

  constructor(private myProjectService: DatabaseProjectServices, private router: Router) {}


  ngOnInit(): void {
    this.myProjectService.getAll().subscribe(projects => {
      console.log(projects)
      this.projects = projects.data
    })
  }

  onClickCompare(): void {
    let items:  string[] = []
    let ids:  number[] = []
    this.checked.forEach((value: number, key: string) =>{
      if (value){
        items.push(key)
        ids.push(value)
      }
    })

    if (items.length > 3){
      this.showErrorMessage("Максимальное число проектов 3")
    }else if (items.length <= 1){
      this.showErrorMessage("Минимальное число проектов для сравнения 2.")
    }else{
      this.router.navigate([`/compare-projects`], {
        queryParams: {
          keys: items,
          value: ids
        }
      });
    }

  }

  childOnChecked(project: CheckedProject){
    if (project.Checked) {
      this.checked.set(project.Name, project.Id)
    }else if (this.checked.has(project.Name)){
      this.checked.delete(project.Name)
    }
    console.log("Parent ", project.Name, project.Id)
  }

  showErrorMessage(msg: string){
    alert(msg)
  }

}
