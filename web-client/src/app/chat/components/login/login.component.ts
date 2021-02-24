import {Component, Inject, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {User} from "../../entities/user";
import {ChatService} from "../../services/chat.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  form: FormGroup;

  constructor(
    @Inject(FormBuilder) fb: FormBuilder,
    private chat: ChatService,
    private router: Router
  ) {
    this.form = fb.group(new User());
    this.chat = chat;
  }

  ngOnInit() {
  }

  login() {
    console.log("logging in")
    this.chat.login(this.form.value);
    this.router.navigate(['/msg']);
  }

}
