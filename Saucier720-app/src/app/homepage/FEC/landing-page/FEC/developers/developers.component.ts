import { Component } from '@angular/core';
import { trigger, state, style, animate, transition } from '@angular/animations';

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
    trigger('slide', [
      transition(':increment', [
        style({ transform: 'translateX(100%)' }),
        animate('0.5s ease-out', style({ transform: 'translateX(0%)' }))
      ]),
      transition(':decrement', [
        style({ transform: 'translateX(-100%)' }),
        animate('0.5s ease-out', style({ transform: 'translateX(0%)' }))
      ])
    ])
  ]
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

  currentProfileIndex: number = 0;

  // Function to navigate to the next developer profile
  nextProfile() {
    this.currentProfileIndex = (this.currentProfileIndex + 1) % this.developerProfiles.length;
  }

  // Function to navigate to the previous developer profile
  previousProfile() {
    this.currentProfileIndex = (this.currentProfileIndex - 1 + this.developerProfiles.length) % this.developerProfiles.length;
  }
}
