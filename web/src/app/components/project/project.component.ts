import {Component, Input} from '@angular/core'
import {IProj} from "../../models/proj.model";
import {ProjectServices} from "../../services/project.services";

@Component({
  selector: 'app-project',
  templateUrl: './project.component.html',
  styleUrls: ['./project.component.css']
})
export class ProjectComponent {
  @Input() project: IProj
  adding = true


  constructor(private projectService: ProjectServices) {
    //TO_DO
    //this.adding = projectService.find(this.project.Name)
  }

  addMyProject(name: String) {
    this.adding = !this.adding
    //TO_DO
  }
}
