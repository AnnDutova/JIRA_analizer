import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {IProj} from "../../models/proj.model";
import {CheckedProject} from "../../models/check-element.model";
import {CheckedSetting} from "../../models/check-setting.model";
import {SettingBox} from "../../models/setting.model";
import {Router} from "@angular/router";
import {DatabaseProjectServices} from "../../services/database-project.services";

@Component({
  selector: 'app-my-project',
  templateUrl: './my-project.component.html',
  styleUrls: ['./my-project.component.css']
})
export class MyProjectComponent implements OnInit{
  @Output() onChecked: EventEmitter<any> = new EventEmitter<{}>();
  @Input() myProject: IProj
  stat: ProjectStat = new ProjectStat()
  processing: boolean
  settings: boolean
  checkboxes: SettingBox[] = []
  setting: Map<any, any> = new Map();

  constructor(private router: Router, private dbProjectService: DatabaseProjectServices) {
  }

  ngOnInit(): void{
    this.processing=false;
    this.settings = false;
    this.checkboxes.push(new SettingBox("Гистограмма, отражающая время, которое задачи провели в открытом состоянии", false, 1 ))
    this.checkboxes.push(new SettingBox("Диаграммы, которые показывают распределение времени по состоянием задач", false, 2 ))
    this.checkboxes.push(new SettingBox("График активности по задачам", false, 3 ))
    this.checkboxes.push(new SettingBox("График сложности задач", false, 4 ))
    this.checkboxes.push(new SettingBox("График, отражающий приоритетность всех задач", false, 5 ))
    this.checkboxes.push(new SettingBox("График, отражающий приоритетность закрытых задач", false, 6 ))



    this.dbProjectService.getProjectStatByID(this.myProject.Id.toString()).subscribe(projects => {

      this.stat.AverageIssuesCount = projects.data["allIssuesCount"]
      this.stat.OpenIssuesCount = projects.data["openIssuesCount"]
      this.stat.AllIssuesCount = projects.data["allIssuesCount"]
      this.stat.AverageTime = projects.data["averageTime"]
      this.stat.CloseIssuesCount = projects.data["closeIssuesCount"]

      console.log(projects.data)
      console.log(this.stat)
    })
  }

  processProject() {
    this.processing = !this.processing
    console.log(this.myProject, this.setting)
    let ids:  number[] = []
    let items = this.myProject.Id

    this.checkboxes.forEach((box: SettingBox) =>{
      if (box.Checked){
        ids.push(Number(box.BoxId))
      }
    })

    this.router.navigate([`/project-stat`], {
      queryParams: {
        keys: items,
        value: ids
      }
    });

  }

  clickOnSettings(){
    this.settings = !this.settings;
  }

  noneSelected(){
    return !this.checkboxes.some(checkbox => checkbox.Checked);
  }

  childOnChecked(setting: CheckedSetting){
    if (setting.Checked) {
      this.setting.set(setting.ProjectName, setting.BoxId)
    }else if (this.setting.has(setting.ProjectName)){
      this.setting.delete(setting.ProjectName)
    }
    this.checkboxes[Number(setting.BoxId) - 1].Checked = setting.Checked
    console.log("Parent ", setting.ProjectName, setting.BoxId, this.checkboxes[Number(setting.BoxId) - 1].Checked)
  }

}

class ProjectStat {
  AllIssuesCount: number;
  AverageIssuesCount: number;
  AverageTime: number;
  CloseIssuesCount: number;
  OpenIssuesCount: number;
}
