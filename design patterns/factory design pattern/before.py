"""
Video exporting app
"""

from abc import ABC, abstractmethod

class VideoExporter(ABC):

    @abstractmethod
    def prepare_export(self, video_data):
        pass

    @abstractmethod
    def do_export(self, folder: str):
        pass

class LosslessVideoExporter(VideoExporter):
    def prepare_export(self, video_data):
        print("Preparing video data for lossless export")

    def do_export(self, folder: str):
        print(f"Exporting video data in lossless format {folder}")

class MediumVideoExporter(VideoExporter):

    def prepare_export(self, video_data):
        print("Preparing video data for medium export")

    def do_export(self, folder: str):
        print(f"Exporting video data in medium format {folder}")

class LowVideoExporter(VideoExporter):

    def prepare_export(self, video_data):
        print("Preparing video data for low export")

    def do_export(self, folder: str):
        print(f"Exporting video data in low format {folder}")


class AudioExporter(ABC):

    @abstractmethod
    def prepare_export(self, video_data):
        pass

    @abstractmethod
    def do_export(self, folder: str):
        pass

class HighAudioExporter(AudioExporter):
    def prepare_export(self, video_data):
        print("Preparing audio data for high export")

    def do_export(self, folder: str):
        print(f"Exporting audio data in high format to {folder}")

class MediumAudioExporter(AudioExporter):
    def prepare_export(self, video_data):
        print("Preparing audio data for medium export")

    def do_export(self, folder: str):
        print(f"Exporting audio data in medium format to {folder}")

class LowAudioExporter(AudioExporter):
    def prepare_export(self, video_data):
        print("Preparing audio data for low export")

    def do_export(self, folder: str):
        print(f"Exporting audio data in low format to {folder}")

def main():
    video_exporter: VideoExporter
    audio_exporter: AudioExporter

    user_input = input("Enter desired output quality (low, high, medium) : ")
    if user_input == "low":
        video_exporter = LowVideoExporter()
        audio_exporter = LowAudioExporter()
    elif user_input == "high":
        video_exporter = LosslessVideoExporter()
        audio_exporter = HighAudioExporter()
    elif user_input == "medium":
        video_exporter = MediumVideoExporter()
        audio_exporter = MediumAudioExporter()
    else:
        raise Exception("Invalid input from user")

    video_exporter.do_export("folder")
    audio_exporter.do_export("folder")

if __name__ == "__main__":
    main()