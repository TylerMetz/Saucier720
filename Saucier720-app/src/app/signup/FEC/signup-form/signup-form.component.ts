import { Component } from '@angular/core';
import { User } from 'src/app/core/interfaces/user';
import { SignupService } from 'src/app/signup.service';

@Component({
  selector: 'app-signup-form',
  templateUrl: './signup-form.component.html',
  styleUrls: ['./signup-form.component.scss']
})
export class SignupFormComponent {
  firstName: string = '';
  lastName: string = '';
  email: string = '';
  userName: string = '';
  password: string = '';

  constructor(private signupService: SignupService) {}

  signup() {
    const user: User = {
      FirstName: this.firstName,
      LastName: this.lastName,
      Email: this.email,
      UserName: this.userName,
      Password: this.password,
    };
    this.signupService.signup(user)
  }
}
