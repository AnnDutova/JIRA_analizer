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
    for (let i = 0; i < this.projects.length; i++) {
      this.dbProjectService.getProjectStatByID(this.ids[i]).subscribe(projects => {
        this.resultReq[i] = projects.data
      })
    }

    var colors = ["blue", "green", "red", "orange", "purple", "black"]

    var openTaskElem = document.getElementById('open-task') as HTMLElement;
    var openTaskTitle = document.getElementById('open-task-title') as HTMLElement;
    this.dbProjectService.getComplitedGraph("1", this.projects).subscribe(info => {
      if (info.data["count"] == null) {
        openTaskElem.remove()
        openTaskTitle.remove()
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
            color: colors[j],
            data: count})
          this.openTaskChart = new Chart(openTaskChartOptions)
        }
      }
    })


    var openStateElem = document.getElementById('open-state') as HTMLElement;
    var resolveStateElem = document.getElementById('resolve-state') as HTMLElement;
    var progressStateElem = document.getElementById('progress-state') as HTMLElement;
    var reopenStateElem = document.getElementById('reopen-state') as HTMLElement;
    var openStateTitle = document.getElementById('open-state-title') as HTMLElement;
    var resolveStateTitle = document.getElementById('resolve-state-title') as HTMLElement;
    var progressStateTitle = document.getElementById('progress-state-title') as HTMLElement;
    var reopenStateTitle = document.getElementById('reopen-state-title') as HTMLElement;
    this.dbProjectService.getComplitedGraph("2", this.projects).subscribe(info => {
      if (info.data["open"] == null) {
        //??????????????, ?????? ?? ??????????-???? ???? ???????????????? ?????????????????????? ???????????????? ????????????
        openStateElem.remove()
        openStateTitle.remove()
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
            type: "spline",
            color: colors[j],
            data: count})
          this.openStateChart = new Chart(openStateChartOptions)
        }
      }
      if (info.data["resolve"] == null) {
        //??????????????, ?????? ?? ??????????-???? ???? ???????????????? ?????????????????????? ???????????????? ????????????
        resolveStateElem.remove()
        resolveStateTitle.remove()
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
            type: "spline",
            color: colors[j],
            data: countResolve})
          this.resolveStateChart = new Chart(resolveStateChartOptions)
        }
      }



      if (info.data["progress"] == null) {
        //??????????????, ?????? ?? ??????????-???? ???? ???????????????? ?????????????????????? ???????????????? ????????????
        progressStateElem.remove()
        progressStateTitle.remove()
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
            type: "spline",
            color: colors[j],
            data: countProgress})
          this.progressStateChart = new Chart(progressStateChartOptions)
        }
      }


      if (info.data["reopen"] == null) {
        //??????????????, ?????? ?? ??????????-???? ???? ???????????????? ?????????????????????? ???????????????? ????????????
        reopenStateElem.remove()
        reopenStateTitle.remove()
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
            type: "spline",
            color: colors[j],
            data: countReopen})
          this.reopenStateChart = new Chart(reopenStateChartOptions)
        }
      }
    })



    var activityByTaskElem = document.getElementById('activity-by-task') as HTMLElement;
    var activityByTaskTitle = document.getElementById('activity-by-task-title') as HTMLElement;
    this.dbProjectService.getComplitedGraph("3", this.projects).subscribe(info => {
      if (info.data["close"] == null) {
        activityByTaskElem.remove()
        activityByTaskTitle.remove()
      }
      else{
        // @ts-ignore
        activityByTaskChartOptions.xAxis["categories"] = info.data["categories"]["all"]

        for (let j = 0; j < this.projects.length; j++){
          var countOpen: any[] = []
          for (let i = 0; i < info.data["categories"]["all"].length; i++) {
            if (info.data["open"][info.data["categories"]["all"][i]] == undefined){
              countOpen.push(0)
            }
            else{
              countOpen.push(info.data["open"][info.data["categories"]["all"][i]][j])
            }
          }
          var countClose: any[] = []
          for (let i = 0; i < info.data["categories"]["all"].length; i++) {
            if (info.data["close"][info.data["categories"]["all"][i]] == undefined){
              countClose.push(0)
            }
            else{
              countClose.push(info.data["close"][info.data["categories"]["all"][i]][j])
            }
          }
          for (let i = 1; i < countOpen.length; i++) {
            if (countOpen[i] == 0){
              countOpen[i] = countOpen[i-1]
            }
          }
          for (let i = 1; i < countClose.length; i++) {
            if (countClose[i] == 0){
              countClose[i] = countClose[i-1]
            }
          }
          activityByTaskChartOptions.series?.push({ name: this.projects[j] + " open",
            type: "spline",
            color: colors[j],
            data: countOpen})
          activityByTaskChartOptions.series?.push({ name: this.projects[j] + " close",
            type: "spline",
            color: colors[j + this.projects.length],
            data: countClose})
          this.activityByTaskChart = new Chart(activityByTaskChartOptions)
        }
      }
    })

    var complexityTaskElem = document.getElementById('complexity-task') as HTMLElement;
    var complexityTaskTitle = document.getElementById('complexity-task-title') as HTMLElement;
    this.dbProjectService.getComplitedGraph("4", this.projects).subscribe(info => {
      if (info.data["categories"] == null) {
        complexityTaskElem.remove()
        complexityTaskTitle.remove()
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
            color: colors[j],
            data: count})
          this.complexityTaskChart = new Chart(complexityTaskChartOptions)
        }
      }
    })

    var taskPriorityElem = document.getElementById('task-priority') as HTMLElement;
    var taskPriorityTitle = document.getElementById('task-priority-title') as HTMLElement;
    this.dbProjectService.getComplitedGraph("5", this.projects).subscribe(info => {
      if (info.data["categories"] == null) {
        taskPriorityElem.remove()
        taskPriorityTitle.remove()
        console.log("Here")
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
            color: colors[j],
            data: count})
          this.taskPriorityChart = new Chart(taskPriorityChartOptions)
        }
      }
    })

    var closeTaskPriorityElem = document.getElementById('close-task-priority') as HTMLElement;
    var closeTaskPriorityTitle = document.getElementById('close-task-priority-title') as HTMLElement;
    this.dbProjectService.getComplitedGraph("6", this.projects).subscribe(info => {
      if (info.data["categories"] == null) {
        closeTaskPriorityElem.remove()
        closeTaskPriorityTitle.remove()
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
            color: colors[j],
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
  resolvedIssuesCount: number;
  reopenedIssuesCount: number;
  progressIssuesCount: number;
}
