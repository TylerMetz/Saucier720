import { Component } from '@angular/core';
import { FormService } from 'src/app/core/services/form/form.service';

@Component({
  selector: 'app-feedback',
  templateUrl: './feedback.component.html',
  styleUrls: ['./feedback.component.scss']
})
export class FeedbackComponent {

  name: string = '';
  email: string = '';
  message: string = '';

  constructor(private formService: FormService) { }

  onSubmit() {
    this.formService.sendFeedback(this.name, this.email, this.message)
      .subscribe(
        () => {
          console.log('Feedback submitted successfully');
          // Clear the form fields
          this.name = '';
          this.email = '';
          this.message = '';
        },
        (error) => {
          console.error('Failed to submit feedback:', error);
        }
      );
  }
}
