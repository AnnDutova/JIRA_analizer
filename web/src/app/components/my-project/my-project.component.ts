import {Component, Input} from '@angular/core';
import {IRequest} from "../../models/request.model";

@Component({
  selector: 'app-my-project',
  templateUrl: './my-project.component.html',
})
export class MyProjectComponent {
  @Input() myProject: IRequest
}
