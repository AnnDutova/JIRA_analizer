import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { HomePageComponent } from './home-page/home-page.component';
import { SetupPageComponent } from './setup-page/setup-page.component';
import {RouterModule} from "@angular/router";
import {ProjectComponent} from "./components/project/project.component";
import {HttpClientModule} from "@angular/common/http";
import { ProjectPageComponent } from './project-page/project-page.component';
import { MyProjectPageComponent } from './my-project-page/my-project-page.component';
import { MyProjectComponent } from './components/my-project/my-project.component';
import {FormsModule} from "@angular/forms";
import {SearchPipe} from "./pipes/search.pipe";
import {NgxPaginationModule} from "ngx-pagination";
import { ComparePageComponent } from './compare-page/compare-page.component';
import {ProjectWithCheckboxComponent} from "./components/checkbox-with-project/checkbox-with-project.component";
import { CompareProjectPageComponent } from './compare-project-page/compare-project-page.component';
import {ChartModule} from "angular-highcharts";

const routes = [
  {path: '', component: HomePageComponent},
  {path: 'setup', component: SetupPageComponent},
  {path: 'projects', component: ProjectPageComponent},
  {path: 'compare', component: ComparePageComponent},
  {path: 'myprojects', component: MyProjectPageComponent},
  {path: 'compare-projects', component: CompareProjectPageComponent}

]

@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    SetupPageComponent,
    ProjectComponent,
    ProjectPageComponent,
    MyProjectComponent,
    MyProjectPageComponent,
    SearchPipe,
    ComparePageComponent,
    ProjectWithCheckboxComponent,
    CompareProjectPageComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot(routes),
    FormsModule,
    NgxPaginationModule,
    ChartModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
