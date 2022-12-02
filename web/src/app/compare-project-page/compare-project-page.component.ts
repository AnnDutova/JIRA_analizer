import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router'
import {DatabaseProjectServices} from "../services/database-project.services";
import {Chart} from "angular-highcharts"
import {openTaskChartOptions} from "./helpers/openTaskChartOptions";
import {taskStateChartOptions} from "./helpers/taskStateChartOptions";

@Component({
  selector: 'app-compare-project-page',
  templateUrl: './compare-project-page.component.html',
  styleUrls: ['./compare-project-page.component.scss']
})

export class CompareProjectPageComponent implements OnInit {
  projects: string[] = []
  ids: string[] = []
  text: ["Общее количество задач", ""]
  resultReq: ReqData[] = []
  taskStateChart = new Chart(taskStateChartOptions)
  openTaskChart = new Chart(openTaskChartOptions)

  constructor(private route: ActivatedRoute, private projectService: DatabaseProjectServices) {
    this.projects = this.route.snapshot.queryParamMap.getAll("keys")
    this.ids = this.route.snapshot.queryParamMap.getAll("value")
  }


  async ngOnInit(): Promise<void> {
    var header = document.getElementById('soft_stat_header') as HTMLElement;
    var sumTask = document.getElementById('summary_task_count') as HTMLElement;
    var taskCountOpen = document.getElementById('open_task_count') as HTMLElement;
    var taskCountClose = document.getElementById('close_task_count') as HTMLElement;
    var averageTaskTime = document.getElementById('average_task_time') as HTMLElement;
    var averageTaskCount = document.getElementById('average_task_count') as HTMLElement;

    for (let i = 0; i < this.projects.length; i++){
      var col = document.createElement('th');
      col.textContent = this.projects[i];
      header.appendChild(col);
      this.projectService.getProjectStatByID(this.ids[i]).subscribe(projects => {
        this.resultReq.push(projects.data)
        var row = document.createElement('td');
        row.textContent = projects.data.allIssuesCount.toString();
        sumTask.appendChild(row);
        row = document.createElement('td');
        row.textContent = projects.data.openIssuesCount.toString();
        taskCountOpen.appendChild(row);
        row = document.createElement('td');
        row.textContent = projects.data.closeIssuesCount.toString();
        taskCountClose.appendChild(row);
        row = document.createElement('td');
        row.textContent = projects.data.averageTime.toString();
        averageTaskTime.appendChild(row);
        row = document.createElement('td');
        row.textContent = projects.data.averageIssuesCount.toString();
        averageTaskCount.appendChild(row);
        console.log("InNotForAwait", projects.data)
      })
    }
  }


}

class ReqData {
  Id: number;
  Key: string;
  Name: string;
  allIssuesCount: number;
  averageIssuesCount: string;
  averageTime: number;
  closeIssuesCount: number;
  openIssuesCount: number;
}
