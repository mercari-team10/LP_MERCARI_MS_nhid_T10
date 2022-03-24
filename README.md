# NHID Microservice

This microservice contains the various methods to manage and control the databses of the patients, doctors and the hospitals.
The various endpoints contained in this microservice are the following :

- `/publicKey` : This endpoint is used to generate the publickey for authentication.

- `/patient/register` : This endpoint is used to register patients.

- `/patient/login` : This endpoint is used to login patients

- `/doctor/register` : This endpoint is used for registration of the doctors.

- `/doctor/login` : This endpoint is used for login of doctors.

- `/doctor/details/:id` : This endpoint is used to retrieve the details of the doctors.

- `/lab/register` : This endpoint is used for adding new labs

- `/lab/login` : This endpoint is used for login of the labs

- `/lab/details/:id` : This endpoint is used for retrieving the lab details

- `/hospital/register` : This endpoint is used for the registration of the labs

- `/hospital/login` : This endpoint is used for login of the hospitals

- `/hospital/details/:id` : This endpoint is used for retrieving the details of the hospitals
