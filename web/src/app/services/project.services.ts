import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {IRequest} from "../models/request.model";

@Injectable({
    providedIn: 'root'
})
export class ProjectServices {
  constructor(private http: HttpClient) {
  }

  getAll(page: number, searchName: String): Observable<IRequest>{
    return this.http.get<IRequest>('http://localhost:8000/api/v1/connector/projects?' +
                                       'limit=10&page='+page + '&search=' + searchName)
  }

  addProject(key: String): Observable<IRequest>{
    return this.http.post<IRequest>('http://localhost:8000/api/v1/connector/updateProject?project='+key, ' ')
  }

  deleteProject(id: Number): Observable<IRequest> {
    return this.http.delete<IRequest>('http://localhost:8000/api/v1/projects/'+id)
  }
}
