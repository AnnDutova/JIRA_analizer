import {Component, Input} from '@angular/core'
import {IProj} from "../../models/proj.model";
import {MyprojectServices} from "../../services/myproject.services";

@Component({
  selector: 'app-project',
  templateUrl: './project.component.html',
  styleUrls: ['./project.component.css']
})
export class ProjectComponent {
  @Input() project: IProj


  constructor(private myprojectService: MyprojectServices) {
  }

  addMyProject(project: IProj) {
    this.myprojectService.addMyProject(project)
  }
}
