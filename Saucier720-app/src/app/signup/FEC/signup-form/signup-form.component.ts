import { Component } from '@angular/core';
import { User } from 'src/app/core/interfaces/user';
import { SignupService } from 'src/app/signup.service';
import { lastValueFrom } from 'rxjs';
import { Router } from '@angular/router';
import { SignupRequest } from 'src/app/core/interfaces/types';

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

  constructor(private signupService: SignupService, private router: Router) {}

  signup(){
    const request: SignupRequest = {
      UserName: this.userName,
      FirstName: this.firstName,
      LastName: this.lastName, 
      Email: this.email,
      Password: this.password
    };
    this.signupService.signup(request).subscribe({
      next: (response: any) => {
        console.log(response);
        this.router.navigate(['/Login']);
      },
      error: (err: any) => {
        console.log(err, 'errors');
      }
    })
  }
}
