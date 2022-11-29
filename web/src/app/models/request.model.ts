import {Links} from "./links.model";

export class IRequest {
  constructor(
  public _links: Links,
  public data: [],
  public message: String,
  public name: String,
  public status: Boolean){}
}


