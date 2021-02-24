import {EventEmitter, Injectable, NgZone, Output} from '@angular/core';
import {WebsocketService} from './web-socket.service';
import {Message} from "../entities/message";
import {User} from "../entities/user";
import {HttpClient} from "@angular/common/http";

/**
 * Best Practice with Observabales
 * @see https://tutorialedge.net/typescript/angular/angular-websockets-tutorial/
 */

const CHAT_URL = 'ws://localhost:7777/ws';

@Injectable({
  providedIn: 'root'
})
export class ChatService {
  public users: User[] = [];
  public me: User;
  @Output() onReady: EventEmitter<User[]> = new EventEmitter();

  constructor(private zone: NgZone, private ws: WebsocketService, private http: HttpClient) {
    this.ws = ws;
    this.ws.connect(CHAT_URL, event => {
      this.zone.run(() => {
        this.subscriber(event);
      })
    });
    this.getUsers()
  }

  private getUsers() {
    this.http.get<User[]>('/users')
      .subscribe(users => this.populateUsers(users['userList']));
  }

  populateUsers(users: User[]) {
    console.log(users);
    for (let user of users) {
      this.users.push(Object.assign(new User(), user));
    }
    this.onReady.emit(this.users);
  }

  login(credential: User) {
    this.ws.send(JSON.stringify(credential));
    this.me = credential;
  }

  sendMessage(user: User, msg: Message) {
    msg.to = user.id;
    user.messages.push(msg);
    this.ws.send(JSON.stringify(msg));
  }

  subscriber(me: MessageEvent) {
    var msg: Message = Object.assign(new Message(), JSON.parse(me.data));
    let user = this.users.find(user => user.id === msg.from);
    if (user) {
      user.messages.push(msg);
    }
  }
}
