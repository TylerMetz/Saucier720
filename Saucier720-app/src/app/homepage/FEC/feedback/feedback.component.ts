import { Component } from '@angular/core';
import { FormService } from 'src/app/core/services/form/form.service';
import { animate, state, style, transition, trigger } from '@angular/animations';
import { post } from 'cypress/types/jquery';

@Component({
  selector: 'app-feedback',
  templateUrl: './feedback.component.html',
  styleUrls: ['./feedback.component.scss'],
  animations: [
    trigger('fadeInOut', [
      state('void', style({ opacity: 0 })),
      transition(':enter, :leave', [
        animate(300) // Adjust the duration (ms) of the fade-in and fade-out here
      ]),
    ]),
  ],
})
export class FeedbackComponent {

  name: string = '';
  email: string = '';
  message: string = '';

  showMessage: boolean = false; // Variable to track if the message is shown
  showForm: boolean = true; // Variable to toggle between form and message

  constructor(private formService: FormService) { }

  onSubmit() {
    if (!this.showMessage) {
      
      // save textbox variables to post variables
      var postName = this.name;
      var postEmail = this.email;
      var postMessage = this.message;

      // reset textbox variables
      this.showForm = false; // Hide the form
      setTimeout(() => {
        this.showMessage = true;
      }, 301); // 301 milliseconds for the animation to finish

      this.name = '';
      this.email = '';
      this.message = '';

      // Send the feedback in POST req
      this.formService.sendFeedback(postName, postEmail, postMessage)
        .subscribe(
          () => {
            
            console.log('Feedback submitted successfully');
          },
          (error) => {
            console.error('Failed to submit feedback:', error);
          }
        );
    } else {
      // Show the form again and reset the message state
      
      this.showMessage = false;
      setTimeout(() => {
        this.showForm = true;
      }, 301); // 301 milliseconds for the animation to finish
    }
  }

}
    
