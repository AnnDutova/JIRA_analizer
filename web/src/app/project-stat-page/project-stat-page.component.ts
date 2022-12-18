import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {DatabaseProjectServices} from "../services/database-project.services";
import {Chart} from "angular-highcharts";
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
  selector: 'project-stat-page',
  templateUrl: './project-stat-page.component.html',
  styleUrls: ['./project-stat-page.component.scss']
})
export class ProjectStatPageComponent implements OnInit {
  projects: string[] = []
  ids: string[] = []
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

    console.log(this.projects, this.ids)

    var openTaskElem = document.getElementById('open-task') as HTMLElement;
    var openTaskTitle = document.getElementById('open-task-title') as HTMLElement;
    if (this.ids.includes("1")){
      this.dbProjectService.getGraph("1", this.projects[0]).subscribe(info => {
        if (info.data == null) {
          openTaskElem.remove()
          openTaskTitle.textContent = "Гистограмма, отражающая время, которое задачи провели в открытом состоянии - " +
            "аналитическая задача недоступна, проект не располагает данными для ее выполнения"
        } else {
          // @ts-ignore
          openTaskChartOptions.xAxis["categories"] = info.data["categories"]
          var count = []
          for (let i = 0; i < info.data["categories"].length; i++) {
            // @ts-ignore
            count.push(info.data["count"][info.data["categories"][i]])
          }
          openTaskChartOptions.series?.push({
            name: this.projects[0],
            type: "column",
            data: count
          })
          this.openTaskChart = new Chart(openTaskChartOptions)
        }
      })
    } else{
      openTaskElem.remove()
      openTaskTitle.remove()
    }



    var openStateElem = document.getElementById('open-state') as HTMLElement;
    var resolveStateElem = document.getElementById('resolve-state') as HTMLElement;
    var progressStateElem = document.getElementById('progress-state') as HTMLElement;
    var reopenStateElem = document.getElementById('reopen-state') as HTMLElement;
    var openStateTitle = document.getElementById('open-state-title') as HTMLElement;
    var resolveStateTitle = document.getElementById('resolve-state-title') as HTMLElement;
    var progressStateTitle = document.getElementById('progress-state-title') as HTMLElement;
    var reopenStateTitle = document.getElementById('reopen-state-title') as HTMLElement;


    this.dbProjectService.getGraph("2", this.projects[0]).subscribe(info => {
      if (!this.ids.includes("2")) {
        openStateElem.remove()
        resolveStateElem.remove()
        progressStateElem.remove()
        reopenStateElem.remove()
        openStateTitle.remove()
        resolveStateTitle.remove()
        progressStateTitle.remove()
        reopenStateTitle.remove()
      }
      else {
        if (info.data["open"] == null) {
          openStateElem.remove()
          openStateTitle.textContent = "Диаграмма, демонстрирующая распределение времени по состоянию \"Open\" - " +
            "аналитическая задача недоступна, проект не располагает данными для ее выполнения"
        } else {
          // @ts-ignore
          openStateChartOptions.xAxis["categories"] = info.data["categories"]["open"]
          var count = []
          for (let i = 0; i < info.data["categories"]["open"].length; i++) {
            // @ts-ignore
            count.push(info.data["open"][info.data["categories"]["open"][i]])
          }
          openStateChartOptions.series?.push({
            name: this.projects[0],
            type: "areaspline",
            data: count
          })
          this.openStateChart = new Chart(openStateChartOptions)
        }
        if (info.data["resolve"] == null) {
          resolveStateElem.remove()
          resolveStateTitle.textContent = "Диаграмма, демонстрирующая распределение времени по состоянию \"Resolve\" - " +
            "аналитическая задача недоступна, проект не располагает данными для ее выполнения"
        } else {
          // @ts-ignore
          resolveStateChartOptions.xAxis["categories"] = info.data["categories"]["resolve"]
          var countResolve = []
          for (let i = 0; i < info.data["categories"]["resolve"].length; i++) {
            // @ts-ignore
            countResolve.push(info.data["resolve"][info.data["categories"]["resolve"][i]])
          }
          resolveStateChartOptions.series?.push({
            name: this.projects[0],
            type: "areaspline",
            data: countResolve
          })
          this.resolveStateChart = new Chart(resolveStateChartOptions)
        }

        if (info.data["progress"] == null) {
          progressStateElem.remove()
          progressStateTitle.textContent = "Диаграмма, демонстрирующая распределение времени по состоянию \"Progress\" - " +
            "аналитическая задача недоступна, проект не располагает данными для ее выполнения"
        } else {
          // @ts-ignore
          progressStateChartOptions.xAxis["categories"] = info.data["categories"]["progress"]
          var countProgress = []
          for (let i = 0; i < info.data["categories"]["progress"].length; i++) {
            // @ts-ignore
            countProgress.push(info.data["progress"][info.data["categories"]["progress"][i]])
          }
          progressStateChartOptions.series?.push({
            name: this.projects[0],
            type: "areaspline",
            data: countProgress
          })
          this.progressStateChart = new Chart(progressStateChartOptions)
        }


        if (info.data["reopen"] == null) {
          reopenStateElem.remove()
          reopenStateTitle.textContent = "Диаграмма, демонстрирующая распределение времени по состоянию \"Reopen\" - " +
            "аналитическая задача недоступна, проект не располагает данными для ее выполнения"
        } else {
          // @ts-ignore
          reopenStateChartOptions.xAxis["categories"] = info.data["categories"]["reopen"]
          var countReopen = []
          for (let i = 0; i < info.data["categories"]["reopen"].length; i++) {
            // @ts-ignore
            countReopen.push(info.data["reopen"][info.data["categories"]["reopen"][i]])
          }
          reopenStateChartOptions.series?.push({
            name: this.projects[0],
            type: "areaspline",
            data: countReopen
          })
          this.reopenStateChart = new Chart(reopenStateChartOptions)
        }
      }
    })


    var activityByTaskElem = document.getElementById('activity-by-task') as HTMLElement;
    var activityByTaskTitle = document.getElementById('activity-by-task-title') as HTMLElement;
    this.dbProjectService.getGraph("3", this.projects[0]).subscribe(info => {
      if (!this.ids.includes("3")) {
        activityByTaskElem.remove()
        activityByTaskTitle.remove()
      }
      else {
        if (info.data["close"] == null) {
          activityByTaskElem.remove()
          activityByTaskTitle.textContent = "График активности по задачам - " +
            "аналитическая задача недоступна, проект не располагает данными для ее выполнения"
        } else {
          // @ts-ignore
          activityByTaskChartOptions.xAxis["categories"] = info.data["categories"]["all"]
          var countOpen = []
          for (let i = 0; i < info.data["categories"]["open"].length; i++) {
            // @ts-ignore
            countOpen.push(info.data["open"][info.data["categories"]["open"][i]])
          }
          console.log(countOpen)
          var countClose = []
          for (let i = 0; i < info.data["categories"]["close"].length; i++) {
            // @ts-ignore
            countClose.push(info.data["close"][info.data["categories"]["close"][i]])
          }
          activityByTaskChartOptions.series?.push({
            name: this.projects[0] + " open",
            type: "areaspline",
            data: countOpen
          })
          activityByTaskChartOptions.series?.push({
            name: this.projects[0] + " close",
            type: "areaspline",
            data: countClose
          })
          this.activityByTaskChart = new Chart(activityByTaskChartOptions)
        }
      }
    })


    var complexityTaskElem = document.getElementById('complexity-task') as HTMLElement;
    var complexityTaskTitle = document.getElementById('complexity-task-title') as HTMLElement;
      this.dbProjectService.getGraph("4", this.projects[0]).subscribe(info => {
        if (!this.ids.includes("4")) {
          complexityTaskElem.remove()
          complexityTaskTitle.remove()
        }
        else {
          if (info.data == null) {
            complexityTaskElem.remove()
            complexityTaskTitle.textContent = "График сложности задач - " +
              "аналитическая задача недоступна, проект не располагает данными для ее выполнения"
          } else {
            // @ts-ignore
            complexityTaskChartOptions.xAxis["categories"] = info.data["categories"]
            var count = []
            for (let i = 0; i < info.data["categories"].length; i++) {
              // @ts-ignore
              count.push(info.data["count"][info.data["categories"][i]])
            }
            complexityTaskChartOptions.series?.push({
              name: this.projects[0],
              type: "column",
              data: count
            })
            this.complexityTaskChart = new Chart(complexityTaskChartOptions)
          }
        }
      })


    var taskPriorityElem = document.getElementById('task-priority') as HTMLElement;
    var taskPriorityTitle = document.getElementById('task-priority-title') as HTMLElement;
      this.dbProjectService.getGraph("5", this.projects[0]).subscribe(info => {
        if (!this.ids.includes("5")) {
          taskPriorityElem.remove()
          taskPriorityTitle.remove()
        }
        else {
          if (info.data["categories"] == null) {
            taskPriorityElem.remove()
            taskPriorityTitle.textContent = "График, отражающий приоритетность всех задач - " +
              "аналитическая задача недоступна, проект не располагает данными для ее выполнения"
          } else {
            // @ts-ignore
            taskPriorityChartOptions.xAxis["categories"] = info.data["categories"]
            var count = []
            for (let i = 0; i < info.data["categories"].length; i++) {
              // @ts-ignore
              count.push(info.data["count"][info.data["categories"][i]])
            }
            taskPriorityChartOptions.series?.push({
              name: this.projects[0],
              type: "column",
              data: count
            })
            this.taskPriorityChart = new Chart(taskPriorityChartOptions)
          }
        }
      })


    var closeTaskPriorityElem = document.getElementById('close-task-priority') as HTMLElement;
    var closeTaskPriorityTitle = document.getElementById('close-task-priority-title') as HTMLElement;
      this.dbProjectService.getGraph("6", this.projects[0]).subscribe(info => {
        if (!this.ids.includes("6")) {
          closeTaskPriorityElem.remove()
          closeTaskPriorityTitle.remove()
        }
        else {
          if (info.data == null) {
            closeTaskPriorityElem.remove()
            closeTaskPriorityTitle.textContent = "График, отражающий приоритетность закрытых задач - " +
              "аналитическая задача недоступна, проект не располагает данными для ее выполнения"
          } else {
            // @ts-ignore
            closeTaskPriorityChartOptions.xAxis["categories"] = info.data["categories"]
            var count = []
            for (let i = 0; i < info.data["categories"].length; i++) {
              // @ts-ignore
              count.push(info.data["count"][info.data["categories"][i]])
            }
            closeTaskPriorityChartOptions.series?.push({
              name: this.projects[0],
              type: "column",
              data: count
            })
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


