# Frame

A site made for enthusiasts for share and watch movies / short films

## Table of Contents 

- [Quick Start](#quick-start)
- [Features](#features)
- [Architucture](#architucture)
- [License](#license)

<a id="quick-start"></a>
## Quick Start

Frame uses docker for easy development, docker and docker-compose are pre-requistes. Follow the following steps once the repository is cloned locally:

**For Backend**
1. Create a `.env` file using example.env as an example
2. Run `docker-compose up --build`

Note: In order to test out if this is working as expected you can `curl http://localhost:8080/ping`, where the result should be `pong`

**For Frontend**
1. `cd frontend`
2. `npm run dev`

<a id="features"></a>
## Features

- [ ] Authentication System: Login, Logout, Register
- [ ] Post Movie: Posting video, which get stored in S3 bucket
- [ ] Watch Movie: Stream video from S3 bucket
- [ ] Movie Rating: Allow users to rate on a particular movie watched on a 1-5 scale
- [ ] Follow System: Abilty of users to follow other users
- [ ] Post comment: Let users share their thoughts on media or share infromation about thier own movies
- [ ] Notifcations: Notify users of events that transpired on the platform
- [ ] Recommendation System: Recomend movies based on other watched movies

<a id="architucture"></a>
## Architucture

**Technologies Used or Will Use:**
- Go and Mux Router (Backend)
- NodeJS and Typescript (Frontend)
- MongoDB(Database)
- Redis (Rate Limiting + Cache)
- Zod (Input validation)
- AWS S3 (Storage)
- TailwindCSS and shadcnUI (Styling)

<a id="license"></a>
## License

Licensed under [MIT License](./LICENSE)
