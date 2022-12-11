import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router'
import {DatabaseProjectServices} from "../services/database-project.services";
import {Chart} from "angular-highcharts"
import {openTaskChartOptions} from "./helpers/openTaskChartOptions";
import {openStateChartOptions} from "./helpers/openStateChartOptions";
import {resolveStateChartOptions} from "./helpers/resolveStateChartOptions";
import {progressStateChartOptions} from "./helpers/progressStateChartOptions";
import {reopenStateChartOptions} from "./helpers/reopenStateChartOptions";
import {activityByTaskChartOptions} from "./helpers/activityByTaskChartOptions";
import {taskPriorityChartOptions} from "./helpers/taskPriorityChartOptions";
import {closeTaskPriorityChartOptions} from "./helpers/closeTaskPriorityChartOptions";
import {complexityTaskChartOptions} from "./helpers/complexityTaskChartOptions";


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
  openTaskChart = new Chart()
  openStateChart = new Chart()
  resolveStateChart = new Chart()
  progressStateChart = new Chart()
  reopenStateChart = new Chart()
  complexityTaskChart = new Chart()
  activityByTaskChart = new Chart()
  taskPriorityChart = new Chart()
  closeTaskPriorityChart = new Chart()

  constructor(private route: ActivatedRoute, private dbProjectService: DatabaseProjectServices) {
    this.projects = this.route.snapshot.queryParamMap.getAll("keys")
    this.ids = this.route.snapshot.queryParamMap.getAll("value")
  }


  ngOnInit(): void {
    var header = document.getElementById('soft_stat_header') as HTMLElement;
    var sumTask = document.getElementById('summary_task_count') as HTMLElement;
    var taskCountOpen = document.getElementById('open_task_count') as HTMLElement;
    var taskCountClose = document.getElementById('close_task_count') as HTMLElement;
    var averageTaskTime = document.getElementById('average_task_time') as HTMLElement;
    var averageTaskCount = document.getElementById('average_task_count') as HTMLElement;

    for (let i = 0; i < this.projects.length; i++) {
      var col = document.createElement('th');
      col.textContent = this.projects[i];
      header.appendChild(col);
      this.dbProjectService.getProjectStatByID(this.ids[i]).subscribe(projects => {
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
      })
    }

    var openTaskElem = document.getElementById('open-task') as HTMLElement;
    this.dbProjectService.getComplitedGraph("1", this.projects).subscribe(info => {
      if (info.data["count"] == null) {
        openTaskElem.remove()
      }
      else{
        // @ts-ignore
        openTaskChartOptions.xAxis["categories"] = info.data["categories"]
        for (let j = 0; j < this.projects.length; j++){
          var count = []
          for (let i = 0; i < info.data["categories"].length; i++){
            // @ts-ignore
            count.push(info.data["count"][info.data["categories"][i]][j])
          }
          openTaskChartOptions.series?.push({ name: this.projects[j],
            type: "column",
            data: count})
          this.openTaskChart = new Chart(openTaskChartOptions)
        }
      }
    })


    var openStateElem = document.getElementById('open-state') as HTMLElement;
    var resolveStateElem = document.getElementById('resolve-state') as HTMLElement;
    var progressStateElem = document.getElementById('progress-state') as HTMLElement;
    var reopenStateElem = document.getElementById('reopen-state') as HTMLElement;
    this.dbProjectService.getComplitedGraph("2", this.projects).subscribe(info => {
      if (info.data["open"] == null) {
        //вывести, что в каком-то из проектов отсутствуют открытые задачи
        openStateElem.remove()
      }
      else{
        // @ts-ignore
        openStateChartOptions.xAxis["categories"] = info.data["categories"]["open"]
        for (let j = 0; j < this.projects.length; j++){
          var count = []
          for (let i = 0; i < info.data["categories"]["open"].length; i++){
            // @ts-ignore
            count.push(info.data["open"][info.data["categories"]["open"][i]][j])
          }
          openStateChartOptions.series?.push({ name: this.projects[j],
            type: "areaspline",
            data: count})
          this.openStateChart = new Chart(openStateChartOptions)
        }
      }
      if (info.data["resolve"] == null) {
        //вывести, что в каком-то из проектов отсутствуют закрытые задачи
        resolveStateElem.remove()
      }
      else{
        // @ts-ignore
        resolveStateChartOptions.xAxis["categories"] = info.data["categories"]["resolve"]
        for (let j = 0; j < this.projects.length; j++){
          var countResolve = []
          for (let i = 0; i < info.data["categories"]["resolve"].length; i++){
            // @ts-ignore
            countResolve.push(info.data["resolve"][info.data["categories"]["resolve"][i]][j])
          }
          resolveStateChartOptions.series?.push({ name: this.projects[j],
            type: "areaspline",
            data: countResolve})
          this.resolveStateChart = new Chart(resolveStateChartOptions)
        }
      }



      if (info.data["progress"] == null) {
        //вывести, что в каком-то из проектов отсутствуют закрытые задачи
        progressStateElem.remove()
      }
      else{
        // @ts-ignore
        progressStateChartOptions.xAxis["categories"] = info.data["categories"]["progress"]
        for (let j = 0; j < this.projects.length; j++){
          var countProgress = []
          for (let i = 0; i < info.data["categories"]["progress"].length; i++){
            // @ts-ignore
            countProgress.push(info.data["progress"][info.data["categories"]["progress"][i]][j])
          }
          progressStateChartOptions.series?.push({ name: this.projects[j],
            type: "areaspline",
            data: countProgress})
          this.progressStateChart = new Chart(progressStateChartOptions)
        }
      }


      if (info.data["reopen"] == null) {
        //вывести, что в каком-то из проектов отсутствуют открытые задачи
        reopenStateElem.remove()
      }
      else{
        // @ts-ignore
        reopenStateChartOptions.xAxis["categories"] = info.data["categories"]["reopen"]
        for (let j = 0; j < this.projects.length; j++){
          var countReopen = []
          for (let i = 0; i < info.data["categories"]["reopen"].length; i++){
            // @ts-ignore
            countReopen.push(info.data["reopen"][info.data["categories"]["reopen"][i]][j])
          }
          reopenStateChartOptions.series?.push({ name: this.projects[j],
            type: "areaspline",
            data: countReopen})
          this.reopenStateChart = new Chart(reopenStateChartOptions)
        }
      }
    })



    var activityByTaskElem = document.getElementById('activity-by-task') as HTMLElement;
    this.dbProjectService.getComplitedGraph("3", this.projects).subscribe(info => {
      if (info.data["close"] == null) {
        activityByTaskElem.remove()
      }
      else{
        // @ts-ignore
        activityByTaskChartOptions.xAxis["categories"] = info.data["categories"]["all"]

        for (let j = 0; j < this.projects.length; j++){
          var countOpen = []
          console.log(info.data)
          for (let i = 0; i < info.data["categories"]["open"].length; i++){
            // @ts-ignore
            console.log(info.data["categories"]["open"][i])
            console.log(info.data["open"][info.data["categories"]["open"][i]])
            console.log(info.data["open"][info.data["categories"]["open"][i]][j])
            countOpen.push(info.data["open"][info.data["categories"]["open"][i]][j])
          }
          var countClose = []
          for (let i = 0; i < info.data["categories"]["close"].length; i++){
            // @ts-ignore
            countClose.push(info.data["close"][info.data["categories"]["close"][i]][j])
          }
          activityByTaskChartOptions.series?.push({ name: this.projects[j] + " open",
            type: "areaspline",
            data: countOpen})
          activityByTaskChartOptions.series?.push({ name: this.projects[j] + " close",
            type: "areaspline",
            data: countClose})
          this.activityByTaskChart = new Chart(activityByTaskChartOptions)
        }
      }
    })

    var complexityTaskElem = document.getElementById('complexity-task') as HTMLElement;
    this.dbProjectService.getComplitedGraph("4", this.projects).subscribe(info => {
      if (info.data["categories"] == null) {
        complexityTaskElem.remove()
      }
      else{
        // @ts-ignore
        complexityTaskChartOptions.xAxis["categories"] = info.data["categories"]
        for (let j = 0; j < this.projects.length; j++){
          var count = []
          for (let i = 0; i < info.data["categories"].length; i++){
            // @ts-ignore
            count.push(info.data["complexity"][info.data["categories"][i]][j])
          }
          complexityTaskChartOptions.series?.push({ name: this.projects[j],
            type: "column",
            data: count})
          this.complexityTaskChart = new Chart(complexityTaskChartOptions)
        }
      }
    })

    var taskPriorityElem = document.getElementById('task-priority') as HTMLElement;
    this.dbProjectService.getComplitedGraph("5", this.projects).subscribe(info => {
      if (info.data["categories"] == null) {
        taskPriorityElem.remove()
      }
      else{
        // @ts-ignore
        taskPriorityChartOptions.xAxis["categories"] = info.data["categories"]
        for (let j = 0; j < this.projects.length; j++){
          var count = []
          for (let i = 0; i < info.data["categories"].length; i++){
            // @ts-ignore
            count.push(info.data["priority"][info.data["categories"][i]][j])
          }
          taskPriorityChartOptions.series?.push({ name: this.projects[j],
            type: "column",
            data: count})
          this.taskPriorityChart = new Chart(taskPriorityChartOptions)
        }
      }
    })

    var closeTaskPriorityElem = document.getElementById('close-task-priority') as HTMLElement;
    this.dbProjectService.getComplitedGraph("6", this.projects).subscribe(info => {
      if (info.data["categories"] == null) {
        closeTaskPriorityElem.remove()
      }
      else{
        // @ts-ignore
        closeTaskPriorityChartOptions.xAxis["categories"] = info.data["categories"]
        for (let j = 0; j < this.projects.length; j++){
          var count = []
          for (let i = 0; i < info.data["categories"].length; i++){
            // @ts-ignore
            count.push(info.data["priority"][info.data["categories"][i]][j])
          }
          closeTaskPriorityChartOptions.series?.push({ name: this.projects[j],
            type: "column",
            data: count})
          this.closeTaskPriorityChart = new Chart(closeTaskPriorityChartOptions)
        }
      }
    })

  }

  ngOnDestroy(): void{
    // @ts-ignore
    openTaskChartOptions.xAxis["categories"] = []
    openTaskChartOptions.series = []
    // @ts-ignore
    openStateChartOptions.xAxis["categories"] = []
    openStateChartOptions.series = []
    // @ts-ignore
    resolveStateChartOptions.xAxis["categories"] = []
    resolveStateChartOptions.series = []
    // @ts-ignore
    progressStateChartOptions.xAxis["categories"] = []
    progressStateChartOptions.series = []
    // @ts-ignore
    reopenStateChartOptions.xAxis["categories"] = []
    reopenStateChartOptions.series = []
    // @ts-ignore
    activityByTaskChartOptions.xAxis["categories"] = []
    activityByTaskChartOptions.series = []
    // @ts-ignore
    taskPriorityChartOptions.xAxis["categories"] = []
    taskPriorityChartOptions.series = []
    // @ts-ignore
    closeTaskPriorityChartOptions.xAxis["categories"] = []
    closeTaskPriorityChartOptions.series = []
    // @ts-ignore
    complexityTaskChartOptions.xAxis["categories"] = []
    complexityTaskChartOptions.series = []
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
