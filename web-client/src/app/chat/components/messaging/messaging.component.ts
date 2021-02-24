import {Component, Input, OnInit} from '@angular/core';
import {User} from "../../entities/user";
import {ChatService} from "../../services/chat.service";
import {Message} from "../../entities/message";

@Component({
  selector: 'app-messaging',
  templateUrl: './messaging.component.html',
  styleUrls: ['./messaging.component.css']
})
export class MessagingComponent implements OnInit {
  public message: string;
  public users: User[];
  public currentUser: User;
  constructor(private chat: ChatService) {}

  ngOnInit() {
    // console.log(this.chat.users);
    this.onChatReady(this.chat.users);
    this.chat.onReady.subscribe(
      users => this.onChatReady(users)
    )
  }

  onChatReady(users: User[]) {
    // console.log(users);
    this.users = users;
    if (this.users) {
      this.currentUser = this.users[0];
    }
  }

  send(user: User) {
    var msg = new Message();
    msg.message = this.message;
    this.chat.sendMessage(user, msg);
  }

  username(id: number) {
    if (id) {
      let user = this.users.find(user => user.id === id);
      if (user) {
        return user.name
      }
    }

    return this.chat.me.name;
  }
}
