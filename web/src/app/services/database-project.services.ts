import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {IRequest} from "../models/request.model";
import {IRequestObject} from "../models/requestObj.model";

@Injectable({
  providedIn: 'root'
})
export class DatabaseProjectServices {
  constructor(private http: HttpClient) {
  }

  getAll(): Observable<IRequest>{
    return this.http.get<IRequest>('http://localhost:8000/api/v1/projects')
  }

  getProjectStatByID(id: string): Observable<IRequestObject> {
    return this.http.get<IRequestObject>('http://localhost:8000/api/v1/projects/'+id)
  }
}
