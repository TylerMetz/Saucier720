import { Component } from '@angular/core';
import { trigger, transition, query, style, animate, group } from '@angular/animations';
const left = [
  query(':enter, :leave', style({ position: 'relative'}), { optional: true }),
  group([
    query(':enter', [style({ transform: 'translateX(-100%)' }), animate('.5s ease-out', style({ transform: 'translateX(0%)' }))], {
      optional: true,
    }),
    query(':leave', [style({ transform: 'translateX(0%)' }), animate('.5s ease-out', style({ transform: 'translateX(150%)' }))], {
      optional: true,
    }),
  ]),
];

const right = [
  query(':enter, :leave', style({ position: 'relative' }), { optional: true }),
  group([
    query(':enter', [style({ transform: 'translateX(150%)' }), animate('.5s ease-out', style({ transform: 'translateX(0%)' }))], {
      optional: true,
    }),
    query(':leave', [style({ transform: 'translateX(0%)' }), animate('.5s ease-out', style({ transform: 'translateX(-150%)' }))], {
      optional: true,
    }),
  ]),
];

interface DeveloperProfile {
  picture: string;
  name: string;
  github: string;
  linkedin: string;
}

@Component({
  selector: 'app-developers',
  templateUrl: './developers.component.html',
  styleUrls: ['./developers.component.scss'],
  animations: [
    trigger('animSlider', [
      transition(':increment', right),
      transition(':decrement', left),
    ]),
  ],
})
export class DevelopersComponent {
  developerProfiles: DeveloperProfile[] = [
    {
      picture: '../../../../assets/images/dev-profile-pictures/riley-dev-profile.png',
      name: 'Riley Cleavenger',
      github: 'https://github.com/rileycleavenger',
      linkedin: 'https://www.linkedin.com/in/rileycleavenger',
    },
    {
      picture: '../../../../assets/images/dev-profile-pictures/sam-dev-profile.jpeg',
      name: 'Sam Forstot',
      github: 'https://github.com/samforstot',
      linkedin: 'https://www.linkedin.com/in/samforstot/',
    },
    {
      picture: '../../../../assets/images/dev-profile-pictures/tyler-dev-profile.jpeg',
      name: 'Tyler Metz',
      github: 'https://github.com/TylerMetz',
      linkedin: 'https://www.linkedin.com/in/tyler-metz-08146b221/',
    },
    {
      picture: '../../../../assets/images/dev-profile-pictures/steele-dev-profile.jpeg',
      name: 'Steele Elliott',
      github: 'https://github.com/steeleelliott03',
      linkedin: 'https://www.linkedin.com/in/evansteeleelliott/',
    },
    {
      picture: '../../../../assets/images/dev-profile-pictures/mike-dev-profile.jpeg',
      name: 'Mike Ciruzzi',
      github: 'https://github.com/mciruzzi1',
      linkedin: 'https://www.linkedin.com/in/mike-ciruzzi/',
    },
    // Add more developer profiles here...
  ];

  counter: number = 0;

  onNext() {
    if (this.counter != this.developerProfiles.length - 1) {
      this.counter++;
    }
  }

  onPrevious() {
    if (this.counter > 0) {
      this.counter--;
    }
  }
}
