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

/**
 * Sends an asynchronous HTTP request to `url`.
 * @param { String } url The request url.
 * @param { Object } data The data to be sent with the request. Can be
 * `null` if param `method` is 'GET'.
 * @param { String } method The HTTP method type.
 * @param { Boolean } withCredentials Should be set to true when using
 * an authorization header such as `Bearer`.
 * @returns The data contained in the response to the request.
 */
const sendRequest = async (url, data = null, method = 'GET', withCredentials = false) => {
  const res = await customAxios({
    url, data, method, withCredentials,
  });
  return res.data;
};

export default sendRequest;
