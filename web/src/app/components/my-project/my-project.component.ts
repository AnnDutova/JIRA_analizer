import {Component, Input, OnInit} from '@angular/core';
import {IProj} from "../../models/proj.model";

@Component({
  selector: 'app-my-project',
  templateUrl: './my-project.component.html',
  styleUrls: ['./my-project.component.css']
})
export class MyProjectComponent implements OnInit{
  @Input() myProject: IProj
  processing: boolean
  settings: boolean
  checkboxes = [
    {name: "box_1", chosen: false},
    {name: "box_2", chosen: false},
    {name: "box_2", chosen: false},
  ]

  ngOnInit(): void{
    this.processing=false;
    this.settings = false;
  }

  processProject() {
    this.processing = !this.processing
  }

  clickOnSettings(){
    this.settings = !this.settings;
  }

  noneSelected(){
    return !this.checkboxes.some(checkbox => checkbox.chosen);
  }
}
