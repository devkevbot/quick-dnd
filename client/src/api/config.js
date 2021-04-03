const { protocol, hostname } = window.location;
const defaultPort = 3000;
const defaultAPIPath = `${protocol}//${hostname}:${defaultPort}`;

export default defaultAPIPath;
