import { Component } from '@angular/core';

interface DeveloperProfile {
  picture: string;
  name: string;
  github: string;
  linkedin: string;
}

@Component({
  selector: 'app-developers',
  templateUrl: './developers.component.html',
  styleUrls: ['./developers.component.scss']
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
      picture: '../../../../assets/images/dev-profile-pictures/tyler-dev-profile.jpeg',
      name: 'Tyler Metz',
      github: 'https://github.com/TylerMetz',
      linkedin: 'https://www.linkedin.com/in/tyler-metz-08146b221/',
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
