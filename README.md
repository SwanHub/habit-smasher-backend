# Bad Habit Smasher (backend)

- The concept is self-evident. The goal of this app is self-improvement. They say the best way to kill an issue is with constant pressure and attention. The Bad Habit Smasher app will serve as a monitoring agent for some bad habits I wish to kick. I hope the app evolves into a general self-improvement device of larger personal significance. 

## Languages / Tools

- `Go`
- `PostgreSQL`

## Deployment 

- `Heroku`

## Endpoint

- `https://habit-smasher.herokuapp.com/`

### 10 Steps to One Victorious Thin Vertical Slice

1. Spinup local `Go` server with hello world endpoint
2. Deploy to `Heroku`
    - `go mod init`
    - `Procfile`
3. Connect `PostgreSQL` DB
4. Create DB Schema + seed
5. Query DB and display data in json
6. Setup CORS, then fetch values from browser
7. Create `React.js` App
8. Fetch and display values from react app
9. Deploy frontend to `Firebase` 
    - create app from console
    - `firebase init` ("build" is public directory)
    - `npm run build`
    - `firebase deploy`
10. Basic styling with `CSS` and `SASS`
    - `npm install node-sass`