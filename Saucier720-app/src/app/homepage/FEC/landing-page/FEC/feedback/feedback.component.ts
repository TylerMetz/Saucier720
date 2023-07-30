import { Component } from '@angular/core';
import { FormService } from 'src/app/core/services/form/form.service';
import { animate, state, style, transition, trigger } from '@angular/animations';

@Component({
  selector: 'app-feedback',
  templateUrl: './feedback.component.html',
  styleUrls: ['./feedback.component.scss'],
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
      this.showMessage = true; // Show the message
      this.showForm = false; // Hide the form
      this.formService.sendFeedback(this.name, this.email, this.message)
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
    
