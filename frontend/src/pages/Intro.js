import React from 'react';
import Box from '@mui/material/Box';
import LoadingButton from '@mui/lab/LoadingButton';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import TextField from '@mui/material/TextField';

import './Intro.css';

export default class Intro extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      email: '',
      form: 'registration',
      isLoading: false,
      password: '',
      submitText: 'Зарегистрироваться',
      tab: 0,
    }

    this.handleEmailInput = this.handleEmailInput.bind(this);
    this.handlePasswordInput = this.handlePasswordInput.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleTabChange = this.handleTabChange.bind(this);
  }

  handleEmailInput(evt) {
    this.setState({email: evt.target.value});
  }

  handlePasswordInput(evt) {
    this.setState({password: evt.target.value});
  }

  handleSubmit(evt) {
    evt.preventDefault();

    const isRegistration = this.state.form === 'registration';

    this.setState({ isLoading: true });

    if (isRegistration) {
      fetch('/api/v1/auth/register',
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: this.state.email,
            password: this.state.password,
          })
        }).then((result) => {
        console.log('REGISTRATION_RESPONSE', result);
        this.setState({ isLoading: false });
      });
    } else {
      fetch('/api/v1/auth/login',
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: this.state.email,
            password: this.state.password,
          })
        }).then((result) => {
        console.log('AUTHORIZATION_RESPONSE', result);
        this.setState({ isLoading: false });
      });
    }
  }

  handleTabChange(evt, tab) {
    const isRegistration = tab === 0;

    this.setState({
      form: isRegistration ? 'registration' : 'authorization',
      submitText: isRegistration ? 'Зарегистрироваться' : 'Войти',
      tab,
    });
  }

  render() {
    return (<div className="intro">
      <div className="text-container">
        <h1 className="title">Фотоклуб</h1>
        <p className="description">Сообщество для фотографов<br/> и моделей</p>
      </div>

      <div className="form-container">
        <Tabs value={this.state.tab} onChange={this.handleTabChange}>
          <Tab label="Регистрация"/>
          <Tab label="Авторизация"/>
        </Tabs>

        <form className="form" onSubmit={this.handleSubmit}>
          <Box my={2}>
            <TextField
              fullWidth
              id="outlined-basic"
              variant="outlined"
              label="Email"
              onInput={this.handleEmailInput}
            />
          </Box>

          <Box my={2}>
            <TextField
              fullWidth
              id="outlined-basic"
              variant="outlined"
              label="Пароль"
              onInput={this.handlePasswordInput}
            />
          </Box>

          <Box mt={4}>
            <LoadingButton
              type="submit"
              variant="contained"
              loading={this.state.isLoading}
            >
              {this.state.submitText}
            </LoadingButton>
          </Box>
        </form>
      </div>
    </div>)
  }
}
