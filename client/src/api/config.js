import axios from 'axios';

const { protocol, hostname } = window.location;
const defaultPort = 3000;
const defaultAPIPath = `${protocol}//${hostname}:${defaultPort}`;

const customAxios = axios.create({
  baseURL: defaultAPIPath,
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json',
  },
});

export default customAxios;
