#!/bin/bash

# Update package list
sudo apt update

# Install ALSA utilities and libraries
sudo apt install -y alsa-utils libasound2 libasound2-plugins

# Install PulseAudio and its ALSA plugin
sudo apt install -y pulseaudio pulseaudio-utils

# Create or update the ALSA configuration file
echo 'pcm.!default {
    type pulse
    fallback "sysdefault"
    hint {
        show on
        description "Default ALSA Output (PulseAudio)"
    }
}

ctl.!default {
    type pulse
    fallback "sysdefault"
}' | sudo tee /etc/asound.conf

echo "ALSA and PulseAudio setup complete!"