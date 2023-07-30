import { Component } from '@angular/core';
import { FormService } from 'src/app/core/services/form/form.service';
import { animate, state, style, transition, trigger } from '@angular/animations';
import { post } from 'cypress/types/jquery';

@Component({
  selector: 'app-feedback',
  templateUrl: './feedback.component.html',
  styleUrls: ['./feedback.component.scss'],
  animations: [
    trigger('formState', [
      state('shown', style({ opacity: 1, transform: 'translateY(0)' })),
      state('hidden', style({ opacity: 0, transform: 'translateY(100%)' })),
      transition('shown <=> hidden', animate('300ms ease-in-out')),
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
      this.showMessage = true; // Show the message
      this.showForm = false; // Hide the form
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
      this.showForm = true;
      this.showMessage = false;
    }
  }

  onResetForm() {
    // Reset the form fields and set the form state to its initial state
    this.name = '';
    this.email = '';
    this.message = '';
    this.showForm = true;
    this.showMessage = false;
  }

}
    
