import { Component } from '@angular/core';
import { trigger, transition, query, style, animate, group } from '@angular/animations';
const left = [
  query(':enter, :leave', style({ position: 'absolute'}), { optional: true }),
  group([
    query(':enter', [style({ transform: 'translateX(-150%)' }), animate('.5s ease-out', style({ transform: 'translateX(0%)' }))], {
      optional: true,
    }),
    query(':leave', [style({ transform: 'translateX(0%)' }), animate('.5s ease-out', style({ transform: 'translateX(150%)' }))], {
      optional: true,
    }),
  ]),
];

const right = [
  query(':enter, :leave', style({ position: 'absolute' }), { optional: true }),
  group([
    query(':enter', [style({ transform: 'translateX(150%)' }), animate('.5s ease-out', style({ transform: 'translateX(0%)' }))], {
      optional: true,
    }),
    query(':leave', [style({ transform: 'translateX(0%)' }), animate('.5s ease-out', style({ transform: 'translateX(-150%)' }))], {
      optional: true,
    }),
  ]),
];

interface FAQ {
  question: string;
  answer: string;
}

@Component({
  selector: 'app-faqs',
  templateUrl: './faqs.component.html',
  styleUrls: ['./faqs.component.scss'],
  animations: [
    trigger('animSlider', [
      transition('0 => 1', right),
      transition('1 => 0', left),
      transition('* => 0', right), // wrapping from last entry to first
      transition('0 => *', left), // wrapping from first entry to last
      transition(':increment', right),
      transition(':decrement', left),
      
    ]),
  ],
})
export class FaqsComponent {
  allFaqs: FAQ[] = [
    {
      question: 'How do I create a user account?',
      answer: 'If this is your first time using the app select the signup button from above, otherwise navigate to the Login page, then to the Sign Up page from there.',
    },
    {
      question: 'How do I login to my account?',
      answer: 'Navigate to the Login page from the homepage, then enter your username and password.',
    },
    {
      question: 'What is the Pantry page?',
      answer: 'The pantry page is where you can view all of your ingredients, and add more as needed.',
    },
    {
      question: 'What is the Deals page?',
      answer: 'The deals page is where you can view all of the current sales at the select grocery stores that we support.',
    },
    {
      question: 'What is the recipes page?',
      answer: 'The recipes page is where you can view all of the recipes that are recommended to you, as well as your personal recipes and favorited recipes.',
    },
    {
      question: 'What is the list page?',
      answer: 'The list page is where you can view and update your shopping list, items you added from recipes or the Deals page will appear there as well.',
    },
    // Add more developer profiles here...
  ];

  counter: number = 0;

  onNext() {
    if (this.counter != this.allFaqs.length - 1) {
      this.counter++;
    } else{
      this.counter = 0;
    }
  }

  onPrevious() {
    if (this.counter > 0) {
      this.counter--;
    } else{
      this.counter = this.allFaqs.length - 1;
    }
  }

}
