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

const routes = [
  {path: '', component: HomePageComponent},
  {path: 'setup', component: SetupPageComponent},
  {path: 'projects', component: ProjectPageComponent},
  {path: 'myprojects', component: MyProjectPageComponent},
]

@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    SetupPageComponent,
    ProjectComponent,
    ProjectPageComponent,
    MyProjectComponent,
    MyProjectPageComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot(routes)
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
