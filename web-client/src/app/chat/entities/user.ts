import {Message} from "./message";

export class User {
  name: string = "";
  id: number;
  password: string = "";
  messages: Message[] = [];
}
