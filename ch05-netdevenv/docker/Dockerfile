# Specify the base image we want to build from.
# This image already has Python 3.8 installed.
FROM python:3.8-slim-buster

# Copy our "requirements.txt" file which contains
# our Python dependencies into the container image,
# and instruct pip to install the packages within.
COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

# Copy our Python script into the image
# This just gets our IP address via ipify.org
# and prints it, as an example.
COPY ./getip.py /getip.py

# Specify the default command to execute when
# this Docker image is run.
CMD [ "python3", "/getip.py"]