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

  ngOnInit(): void{
    this.processing=false;
    this.settings = false;
  }

  processProject(name: String) {
   // this.processing = !this.processing
  }

  clickOnSettings(){
    this.settings = !this.settings;
  }
}
